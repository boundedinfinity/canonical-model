import _ from "lodash";
import { stringUtils as su } from "./utils.ts";

export type CaseType =
  | "phrase"
  | "pascal"
  | "camel"
  | "snake"
  | "snake-lower"
  | "snake-upper"
  | "kebab"
  | "kebab-lower"
  | "kebab-upper";

export class Caser {
  options: {
    safeChar: string;
    ignoreSymbols: string[];
    compactSpaces: boolean;
  } = {
    safeChar: " ",
    compactSpaces: true,
    ignoreSymbols: [],
  };

  constructor(options?: Partial<typeof this.options>) {
    _.merge(this.options, options);
  }

  convert(s: string, from: CaseType, to: CaseType): string {
    let result: string;

    switch (from) {
      case "phrase":
        result = s;
        break;
      default:
        throw new Error(`invalid from conversion: ${from}`);
    }

    switch (to) {
      case "phrase":
        // Don't need to do anything since already in phase form
        break;
      case "camel":
        result = this.phrase2Camel(result);
        break;
      case "kebab":
        result = this.phrase2Kebab(result);
        break;
      case "kebab-upper":
        result = this.phrase2KebabUpper(result);
        break;
      case "kebab-lower":
        result = this.phrase2KebabLower(result);
        break;
      case "snake":
        result = this.phrase2Snake(result);
        break;
      case "snake-upper":
        result = this.phrase2SnakeUpper(result);
        break;
      case "snake-lower":
        result = this.phrase2SnakeLower(result);
        break;
      default:
        throw new Error(`invalid to conversion: ${to}`);
    }

    return result;
  }

  private symbols(options: typeof this.options): string[] {
    const results = su.symbols.filter((s) =>
      !options.ignoreSymbols.includes(s)
    );
    return results;
  }

  private mergeOptions(
    options?: Partial<typeof this.options>,
  ): typeof this.options {
    const result = _.cloneDeep(this.options);
    _.merge(result, options);
    return result;
  }

  kebab2phrase(s: string, options?: Partial<typeof this.options>): string {
    _.merge(this.options, options);
    return su.compactSpaces(s).split("-").join(" ");
  }

  snake2phrase(s: string, options?: Partial<typeof this.options>): string {
    _.merge(this.options, options);
    return su.compactSpaces(s).split("_").join(" ");
  }

  camel2phrase(s: string, options?: Partial<typeof this.options>): string {
    _.merge(this.options, options);
    const comps: string[] = [];
    let wordStart = 0;

    for (let i = 0; i < s.length; i++) {
      const c = s[i];

      if (su.isUpperCase(c)) {
        comps.push(s.substring(wordStart, i));
        wordStart = i;
      }
    }

    comps.push(s.slice(wordStart));

    return comps.join(" ").trim();
  }

  pascal2phrase(s: string, options?: Partial<typeof this.options>): string {
    return this.camel2phrase(s, options);
  }

  phrase2Kebab(s: string, options?: Partial<typeof this.options>): string {
    const roptions = this.mergeOptions(options);
    const rsymbols = this.symbols(roptions);

    return su.pipe(
      s,
      su.replaceWithFn(rsymbols, ""),
      su.compactSpaces,
      su.replaceWithFn([" "], "_"),
    );
  }

  phrase2KebabUpper(s: string, options?: Partial<typeof this.options>): string {
    return this.phrase2Kebab(s, options).toLocaleUpperCase();
  }

  phrase2KebabLower(s: string, options?: Partial<typeof this.options>): string {
    return this.phrase2Kebab(s, options).toLocaleLowerCase();
  }

  phrase2Snake(s: string, options?: Partial<typeof this.options>): string {
    const roptions = this.mergeOptions(options);
    const rsymbols = this.symbols(roptions);

    return su.pipe(
      s,
      su.replaceWithFn(rsymbols, ""),
      su.compactSpaces,
      su.replaceWithFn([" "], "_"),
    );
  }

  phrase2SnakeUpper(s: string, options?: Partial<typeof this.options>): string {
    return this.phrase2Snake(s, options).toLocaleUpperCase();
  }

  phrase2SnakeLower(s: string, options?: Partial<typeof this.options>): string {
    return this.phrase2Snake(s, options).toLocaleLowerCase();
  }

  phrase2Camel(s: string, options?: Partial<typeof this.options>): string {
    return su.lowerCaseFirst(this.phrase2Pascal(s, options));
  }

  phrase2Pascal(s: string, options?: Partial<typeof this.options>): string {
    const roptions = this.mergeOptions(options);
    const rsymbols = this.symbols(roptions);

    return su.pipe(
      s,
      su.replaceWithFn(rsymbols, ""),
      su.compactSpaces,
      su.lowerCase,
      su.splitThen(" ", su.uppperCaseFirst),
      su.replaceWithFn([" "], ""),
    );
  }
}
