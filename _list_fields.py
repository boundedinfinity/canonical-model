import re
text = open('go/idiomatic/specification/iso/iso3166/country.go').read()
fields = re.findall(r'^\t([A-Z]\w+)\s+Country\s*$', text, re.MULTILINE)
print('\n'.join(fields))
