package com.donatasd.duplicates;

import java.util.HashMap;

public class CountingDuplicates {

    public static int duplicateCount(String text) {
        var map = new HashMap<Character, Integer>();
        var count = 0;
        for (Character character : text.toLowerCase().toCharArray()) {
            var occurrences = map.getOrDefault(character, 0);
            if (occurrences == 1) {
                count++;
            }
            map.put(character, occurrences + 1);
        }
        return count;
    }
}

