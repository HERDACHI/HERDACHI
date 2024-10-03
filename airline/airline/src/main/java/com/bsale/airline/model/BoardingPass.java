package com.bsale.airline.model;

import javax.persistence.*;

@Entity
@Table(name = "boarding_pass")
public class BoardingPass {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "boarding_pass_id")
    private Integer boardingPassId;

    @Column(name = "purchase_id")
    private Integer purchaseId;

    @Column(name = "passenger_id")
    private Integer passengerId;

    @Column(name = "seat_type_id")
    private Integer seatTypeId;

    @Column(name = "seat_id")
    private Integer seatId;

    //@ManyToOne(fetch = FetchType.LAZY)
    @ManyToOne
    @JoinColumn(name = "flight_id")
    private Flight flight;

    @ManyToOne
    @JoinColumn(name = "passenger_id", insertable = false, updatable = false)
    private Passenger passenger;

// Getters

    public Integer getBoardingPassId() {
        return boardingPassId;
    }

    public Integer getPurchaseId() {
        return purchaseId;
    }

    public Integer getPassengerId() {
        return passengerId;
    }

    public Integer getSeatTypeId() {
        return seatTypeId;
    }

    public Integer getSeatId() {
        return seatId;
    }

    public String getName() {
        return passenger.getName();
    }

    public String getCountry() {
        return passenger.getCountry();
    }

    public Integer getAge() {
        return passenger.getAge();
    }

    public String getDni() {
        return passenger.getDni();
    }
}