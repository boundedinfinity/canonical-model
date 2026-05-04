import re, sys, os, html

base = os.path.dirname(os.path.abspath(__file__))
country_go   = open(os.path.join(base, 'go/idiomatic/specification/iso/iso3166/country.go')).read()
countries_go = open(os.path.join(base, 'go/idiomatic/specification/iso/iso3166/countries.go')).read()

# ── collect every named Country field from the struct ──────────────────────────
field_names = re.findall(r'^\t([A-Z]\w+)\s+Country\s*$', country_go, re.MULTILINE)
print(f"field count: {len(field_names)}", file=sys.stderr)

def norm(s):
    s = html.unescape(s).lower()
    for fr, to in [
        ('à','a'),('á','a'),('â','a'),('ã','a'),('ä','a'),('å','a'),('ā','a'),('ă','a'),('ą','a'),
        ('ç','c'),('ć','c'),('č','c'),('è','e'),('é','e'),('ê','e'),('ë','e'),('ē','e'),('ė','e'),('ę','e'),
        ('ì','i'),('í','i'),('î','i'),('ï','i'),('ī','i'),('ñ','n'),('ń','n'),
        ('ò','o'),('ó','o'),('ô','o'),('õ','o'),('ö','o'),('ø','o'),('ō','o'),
        ('ù','u'),('ú','u'),('û','u'),('ü','u'),('ū','u'),('ý','y'),('ÿ','y'),
        ('æ','ae'),('œ','oe'),('đ','d'),('ð','d'),('ł','l'),('ß','ss'),
    ]:
        s = s.replace(fr, to)
    return re.sub(r'[^a-z0-9]', '', s)

field_map = {norm(f): f for f in field_names}

# Explicit overrides for countries whose formal ISO name differs from the Go field name.
# Key = exact Name string as it appears in countries.go (HTML entities may be present).
EXPLICIT = {
    # HTML-entity apostrophe variants
    "Côte d&#39;Ivoire":                         "CoteDivoire",
    "Lao People&#39;s Democratic Republic":      "Laos",
    "Korea, Democratic People&#39;s Republic of": None,  # no dedicated field
    # Formal/long name → short field name
    "Iran, Islamic Republic of":                 "Iran",
    "Micronesia, Federated States of":           "Micronesia",
    "Moldova, Republic of":                      "Moldova",
    "Netherlands, Kingdom of the":              "Netherlands",
    "Palestine, State of":                       "PalestineState",
    "Russian Federation":                        "Russia",
    "Syrian Arab Republic":                      "Syria",
    "Taiwan, Province of China":                 "Taiwan",
    "Tanzania, United Republic of":              "Tanzania",
    "Türkiye":                                   "Turkey",
    "United Kingdom of Great Britain and Northern Ireland": "UnitedKingdom",
    "United States of America":                  "UnitedStates",
    "Venezuela, Bolivarian Republic of":         "Venezuela",
    "Korea, Republic of":                        None,  # no dedicated field
    # Bonaire entry uses multi-island name
    "Bonaire, Sint Eustatius and Saba":          "Bonaire",
}

# ── find the all: []Country{ ... } block ──────────────────────────────────────
all_start = countries_go.index('\tall: []Country{')
pos = all_start + len('\tall: []Country{')
depth = 1
while depth > 0:
    c = countries_go[pos]
    if c == '{': depth += 1
    elif c == '}': depth -= 1
    pos += 1
all_end = pos  # points just past the closing }
print(f"all block chars: {all_start}–{all_end}", file=sys.stderr)

# ── extract individual country entries ────────────────────────────────────────
inner = countries_go[all_start + len('\tall: []Country{') : all_end - 1]
entries = []
i = 0
while i < len(inner):
    if inner[i] == '{':
        depth = 1; j = i + 1
        while j < len(inner) and depth > 0:
            if inner[j] == '{': depth += 1
            elif inner[j] == '}': depth -= 1
            j += 1
        entries.append(inner[i:j])
        i = j
    else:
        i += 1

print(f"entries in all: {len(entries)}", file=sys.stderr)

# ── map each entry to its field name ─────────────────────────────────────────
no_field = []     # entries that stay in all (no struct field exists)
mapped   = []     # (field_name, entry_text)  — go to named fields

for entry in entries:
    m = re.search(r'Name:\s+"([^"]+)"', entry)
    if not m:
        no_field.append(entry)
        continue
    name = m.group(1)

    # 1) Try explicit override
    if name in EXPLICIT:
        field = EXPLICIT[name]
        if field is None:
            no_field.append(entry)
        else:
            mapped.append((field, entry))
        continue

    # 2) Try normalised automatic match
    key = norm(name)
    if key in field_map:
        mapped.append((field_map[key], entry))
    else:
        no_field.append(entry)

print(f"mapped: {len(mapped)}, staying in all: {len(no_field)}", file=sys.stderr)

# ── build named field blocks ──────────────────────────────────────────────────
def entry_to_named_field(field_name, entry):
    # entry is "{\n\t\t\tName: ...\n\t\t}" (3-tab-indented body)
    body = entry[1:-1]  # strip outer { }
    lines = body.split('\n')
    reindented = []
    for line in lines:
        if line.startswith('\t\t\t'):
            reindented.append('\t\t' + line[3:])
        else:
            reindented.append(line)
    inner_body = '\n'.join(reindented)
    return f'\t{field_name}: Country{{{inner_body}\t}},'

named_blocks = [entry_to_named_field(f, e) for f, e in mapped]
new_named_text = '\n'.join(named_blocks)

# ── build the new all: block with only no-field entries ──────────────────────
if no_field:
    remaining = '\n'.join(no_field)
    new_all_block = f'\tall: []Country{{\n\t\t{remaining.strip()}\n\t}},'
else:
    new_all_block = ''

# ── splice into the file ──────────────────────────────────────────────────────
# Find the trailing comma + newline after all_end
tail_end = all_end
if countries_go[tail_end:tail_end+2] == ',\n':
    tail_end += 2
elif countries_go[tail_end] == ',':
    tail_end += 1

if new_all_block:
    replacement = new_named_text + '\n' + new_all_block + '\n'
else:
    replacement = new_named_text + '\n'

new_content = countries_go[:all_start] + replacement + countries_go[tail_end:]

# ── replace the init() block to rebuild all from named fields + remaining ─────
old_init_block_re = re.compile(
    r'\nfunc init\(\) \{.*?\n\}\n\nfunc populateNamedCountriesFromAll\(\) \{.*?\n\}\n\nfunc countryKey\(.*?\n\}\n',
    re.DOTALL,
)

new_init = r'''
func init() {
	v := reflect.ValueOf(Countries)
	t := v.Type()
	countryType := reflect.TypeOf(Country{})
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type == countryType {
			c := v.Field(i).Interface().(Country)
			if c.Name != "" {
				Countries.all = append(Countries.all, c)
			}
		}
	}
}

'''

new_content = old_init_block_re.sub(new_init, new_content)

# ── fix imports ───────────────────────────────────────────────────────────────
new_content = new_content.replace(
    'import (\n\t"reflect"\n\t"strings"\n\t"unicode"\n)',
    'import (\n\t"reflect"\n)',
)

out_path = os.path.join(base, 'go/idiomatic/specification/iso/iso3166/countries.go')
open(out_path, 'w').write(new_content)
print("Done. File written.", file=sys.stderr)

