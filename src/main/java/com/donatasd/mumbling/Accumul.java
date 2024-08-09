package com.donatasd.mumbling;

public class Accumul {

    public static String accum(String s) {
        if (s.isEmpty()) {
            return s;
        }
        StringBuilder result = new StringBuilder(String.valueOf(s.charAt(0)));
        for (int i = 1; i < s.length(); i++) {
            var str = String.valueOf(s.charAt(i));
            result.append("-").append(str.toUpperCase()).append(String.valueOf(s.charAt(i)).toLowerCase().repeat(i));
        }
        return result.toString();
    }
}