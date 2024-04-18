'use strict';

const wordCharacteristics = (words) => {
    let wordMap = {},
        uniqueWords = []

    for (let i = 0; i < words.length; i++) {
        if (wordMap[words[i]] === undefined) {
            wordMap[words[i]] = 1
            uniqueWords.push(words[i])
        } else {
            wordMap[words[i]] += 1
        }
    }

    return [wordMap, words.length, uniqueWords]
}

const checkCharacters = (firstMap, firstLen, firstUnique, secondMap, secondLen, secondUnique) => {
    if (firstLen != secondLen) {
        return false
    }

    if (firstUnique.length != secondUnique.length) {
        return false
    }

    for (const prop in firstMap) {
        if (firstMap[prop] !== secondMap[prop]) {
            return false
        }
    }

    return true
}

const findPattern = (input, patterns) => {
    const [patternMap, patternLen, uniquePattern] = wordCharacteristics(patterns)
    let sameIndexes = [],
        samePattern = [];

    for (let i = 0; i < input.length; i++) {
        let words = ''
        for (let j = 0; j < patternLen && (j+i) < input.length; j++) {
            words = words + input[i+j]
        }

        const [wordMap, wordLen, uniqueWord] = wordCharacteristics(words)

        if (checkCharacters(patternMap, patternLen, uniquePattern, wordMap, wordLen, uniqueWord)) {
            sameIndexes.push(i)
            samePattern.push(words)
        }
    }

    return sameIndexes.length
}

console.log(findPattern("ABDCKDHJABDCBDAUOQJDBADCLDLCHBCBABCBAABCDAJDBABDCABDABDBCADBCASSJGABCDAUTACBDBQWUDNCDBCADKDHABDJGBDABCBDBADCACADBADBCBAD", "ABCD"))