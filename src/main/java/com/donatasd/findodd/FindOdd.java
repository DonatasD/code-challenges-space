package com.donatasd.findodd;

import java.util.HashMap;

public class FindOdd {
    public static int findIt(int[] numbers) {
        var map = new HashMap<Integer, Boolean>();
        for (int number: numbers) {
            var value = map.getOrDefault(number, true);
            map.put(number, !value);
        }
        return map.entrySet().stream().filter((entry) -> !entry.getValue()).findFirst().get().getKey();
    }
}
