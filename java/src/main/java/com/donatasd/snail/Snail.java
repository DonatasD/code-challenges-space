package com.donatasd.snail;

public class Snail {

    public enum Flow {
        Right,
        Bottom,
        Left,
        Top
    }

    public static int[] snail(int[][] array) {
        var result = new int[array.length * array[0].length];
        var x = 0;
        var y = 0;
        var flow = Flow.Right;
        var xLimit = array.length - 1;
        var yLimit = array[0].length - 1;
        var xStart = 0;
        var yStart = 0;
        for (int i = 0; i < result.length; i++) {
            result[i] =  array[y][x];
            switch (flow) {
                case Right:
                    x++;
                    if (x == xLimit) {
                        flow = Flow.Bottom;
                        xLimit--;
                    }
                    break;
                case Bottom:
                    y++;
                    if (y == yLimit) {
                        flow = Flow.Left;
                        yLimit--;
                    }
                    break;
                case Left:
                    x--;
                    if (x == xStart) {
                        flow = Flow.Top;
                        yStart++;
                    }
                    break;
                case Top:
                    y--;
                    if (y == yStart) {
                        flow = Flow.Right;
                        xStart++;
                    }
                    break;
            }
        }
        return result;
    }
}