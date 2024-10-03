package com.bsale.airline.model;

import javax.persistence.Entity;
import javax.persistence.*;
import java.util.List;

@Entity
@Table(name = "passenger")
public class Passenger {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "passenger_id")
    private Long passengerId;

    @Column(name = "dni")
    private String dni;

    @Column(name = "name")
    private String name;

    @Column(name = "age")
    private int age;

    @Column(name = "country")
    private String country;

    @OneToMany(mappedBy = "passenger")
    private List<BoardingPass> passengerList;

//Getters


    public Long getPassengerId() {
        return passengerId;
    }

    public String getDni() {
        return dni;
    }

    public String getName() {
        return name;
    }

    public int getAge() {
        return age;
    }

    public String getCountry() {
        return country;
    }

}