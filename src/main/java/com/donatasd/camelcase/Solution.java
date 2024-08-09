package com.donatasd.camelcase;

import java.util.List;

class Solution {

    static String toCamelCase(String s) {
        var charsToReplace = List.of('-', '_');
        var builder = new StringBuilder();

        for (int i = 0; i < s.length(); i++) {
            var character = Character.valueOf(s.charAt(i));
            if (charsToReplace.contains(character) && i != 0) {
                builder.append(String.valueOf(s.charAt(i + 1)).toUpperCase());
                i++;
            } else {
                builder.append(character);
            }
        }
        return builder.toString();
    }
}