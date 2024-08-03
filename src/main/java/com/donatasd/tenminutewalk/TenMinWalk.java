package com.donatasd.tenminutewalk;

public class TenMinWalk {

    private static final Integer WALK_DURATION = 10;
    public static boolean isValid(char[] walk) {
        if (walk.length != WALK_DURATION) {
            return false;
        }
        var startingX = 0;
        var startingY = 0;
        for (char step: walk) {
            switch (step) {
                case 'n':
                    startingY++;
                    break;
                case 's':
                    startingY--;
                    break;
                case 'w':
                    startingX--;
                    break;
                case 'e':
                    startingX++;
                    break;
                default:
                    return false;
            }
        }
        return startingX == 0 && startingY == 0;
    }
}