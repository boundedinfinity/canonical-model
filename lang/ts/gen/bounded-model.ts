type Prettify<T> =
  & {
    [K in keyof T]: T[K];
    // deno-lint-ignore ban-types
  }
  & {};

export type BoundedType = Prettify<
  | BoundedObject
  | BoundedRef
  | BoundedArray
  | BoundedString
  | BoundedNumber
  | BoundedInteger
  | BoundedBoolean
>;

export interface BoundedObject {
  kind: "object";
  parent?: BoundedType;
  name?: string;
  properties: BoundedType[];
  optional?: boolean;
  oneOf?: string[];
}

export interface BoundedRef {
  kind: "ref";
  parent?: BoundedType;
  name?: string;
  optional?: boolean;
  ref: string;
}

export interface BoundedArray {
  kind: "array";
  parent?: BoundedType;
  name?: string;
  items: BoundedType;
  optional?: boolean;
  min?: number;
  max?: number;
}

export interface BoundedString {
  kind: "string";
  parent?: BoundedType;
  name?: string;
  optional?: boolean;
  unique?: boolean;
  primaryKey?: boolean;
  indexed?: boolean;
  min?: number;
  max?: number;
  anyOf?: string[];
  oneOf?: string[];
  allOf?: string[];
  noneOf?: string[];
  startsWith?: string;
  endsWith?: string;
  matches?: string;
  searchable?: boolean;
}

export interface BoundedInteger {
  kind: "integer";
  parent?: BoundedType;
  name?: string;
  optional?: boolean;
  primaryKey?: boolean;
  unique?: boolean;
  indexed?: boolean;
  min?: number;
  max?: number;
  anyOf?: number[];
  allOf?: number[];
  noneOf?: number[];
  multipleOf?: number;
  searchable?: boolean;
}

export interface BoundedNumber {
  kind: "number";
  parent?: BoundedType;
  name?: string;
  optional?: boolean;
  primaryKey?: boolean;
  unique?: boolean;
  indexed?: boolean;
  min?: number;
  max?: number;
  anyOf?: number[];
  allOf?: number[];
  noneOf?: number[];
  multipleOf?: number;
  searchable?: boolean;
}

export interface BoundedBoolean {
  kind: "boolean";
  parent?: BoundedType;
  name?: string;
  optional?: boolean;
  searchable?: boolean;
}
