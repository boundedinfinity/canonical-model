import { assertEquals } from "@std/assert";
import { Caser } from './caser.ts'

const phrase1 = " a'    Phrase wiTh Other WoRDs"
const phrase2 = " 1'    Phrase wiTh Other WoRDs"

Deno.test('pascal to phrase', () => {
    const caser = new Caser()
    assertEquals(caser.pascal2phrase("APhraseWithOtherWords"), 'a Phrase With Other Words')
    assertEquals(caser.pascal2phrase(" "), '')
    assertEquals(caser.pascal2phrase("1PhraseWithOtherWords"), '1 Phrase With Other Words')
})


Deno.test('camel to phrase', () => {
    const caser = new Caser()
    assertEquals(caser.camel2phrase("aPhraseWithOtherWords"), 'A Phrase With Other Words')
    assertEquals(caser.camel2phrase(" "), '')
    assertEquals(caser.camel2phrase("1PhraseWithOtherWords"), '1 Phrase With Other Words')
})

Deno.test('phrase to pascal', () => {
    const caser = new Caser()
    assertEquals(caser.phrase2Pascal(phrase1), 'APhraseWithOtherWords')
    assertEquals(caser.phrase2Pascal(" "), '')
    assertEquals(caser.phrase2Pascal(phrase2), '1PhraseWithOtherWords')
})

Deno.test('phrase to camel', () => {
    const caser = new Caser()
    assertEquals(caser.phrase2Camel(phrase1), 'aPhraseWithOtherWords')
    assertEquals(caser.phrase2Camel(" "), '')
    assertEquals(caser.phrase2Camel(phrase2), '1PhraseWithOtherWords')
})

Deno.test('phrase to snake', () => {
    const caser = new Caser()
    assertEquals(caser.phrase2Snake(phrase1), 'a_Phrase_wiTh_Other_WoRDs')
    assertEquals(caser.phrase2Snake(" "), '')
    assertEquals(caser.phrase2Snake(phrase2), '1_Phrase_wiTh_Other_WoRDs')

    assertEquals(caser.phrase2SnakeLower(phrase1), 'a_phrase_with_other_words')
    assertEquals(caser.phrase2SnakeLower(" "), '')
    assertEquals(caser.phrase2SnakeLower(phrase2), '1_phrase_with_other_words')
})
