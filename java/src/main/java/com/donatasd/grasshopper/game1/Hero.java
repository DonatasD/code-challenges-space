package com.donatasd.grasshopper.game1;

public class Hero {
    String name;
    String position = "00";
    Integer health = 100;
    Integer damage = 5;
    Integer experience = 0;

    public Hero() {
        this.name = "Hero";
    }

    public Hero(String name) {
        this.name = name;
    }
}