package com.donatasd.stripcomments;

import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class StripComments {

    public static String stripComments(String text, String[] commentSymbols) {
        var symbolsToEscape = List.of("]", "[");
        var regex = new StringBuilder();
        regex.append("(\\s*[");
        for (String symbol: commentSymbols) {
            if (symbolsToEscape.contains(symbol)) {
                regex.append("\\" + "symbol");
            } else {
                regex.append(symbol);
            }
        }
        regex.append("].+)|((?m)^\\\\s+$)");
        final Pattern pattern = Pattern.compile(regex.toString(), Pattern.MULTILINE);
        final Matcher matcher = pattern.matcher(text);
        return matcher.replaceAll("");
    }
}
