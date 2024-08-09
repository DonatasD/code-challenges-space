package com.donatasd.permutations;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

class Permutations {

    public static List<String> singlePermutations(String value) {
        Set<String> uniqueValues = new HashSet<>();
        if (value.length() == 1) {
            uniqueValues.add(value);
        } else {
            for (int i = 0; i < value.length(); i++) {
                List<String> temp = singlePermutations(value.substring(0, i) + value.substring(i + 1));
                for (String string : temp) {
                    uniqueValues.add(value.charAt(i) + string);
                }
            }
        }

        return new ArrayList<>(uniqueValues);
    }
}