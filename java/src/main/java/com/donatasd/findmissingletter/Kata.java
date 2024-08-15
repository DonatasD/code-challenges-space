package com.donatasd.findmissingletter;

public class Kata
{
    public static char findMissingLetter(char[] array) {
        char result = 0;
        for (int i = 0; i < array.length - 1; i++) {
            if (array[i + 1] - array[i] != 1) {
                result = (char) ((int) array[i + 1] - 1);
                break;
            }
        }
        return result;
    }
}
