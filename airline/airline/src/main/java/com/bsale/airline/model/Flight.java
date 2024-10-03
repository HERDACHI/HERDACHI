package com.bsale.airline.model;

import javax.persistence.*;
import java.util.List;

@Entity
@Table(name = "flight")
public class Flight {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "flight_id")
    private Integer flightId;

    @Column(name = "takeoff_date_time")
    private Long takeoffDateTime;

    @Column(name = "takeoff_airport")
    private String takeoffAirport;

    @Column(name = "landing_date_time")
    private Long landingDateTime;

    @Column(name = "landing_airport")
    private String landingAirport;

    @Column(name = "airplane_id")
    private Integer airplaneId;

    //@OneToMany(fetch =FetchType.LAZY ,mappedBy="flight", cascade = CascadeType.ALL)
    @OneToMany(mappedBy = "flight")
    private List<BoardingPass> BoardingPassList;

// Getters

    public Integer getFlightId() {
        return flightId;
    }

    public Long getTakeoffDateTime() {
        return takeoffDateTime;
    }

    public String getTakeoffAirport() {
        return takeoffAirport;
    }

    public Long getLandingDateTime() {
        return landingDateTime;
    }

    public String getLandingAirport() {
        return landingAirport;
    }

    public Integer getAirplaneId() {
        return airplaneId;
    }

    public List<BoardingPass> getPassengers() {
        return BoardingPassList;
    }

}