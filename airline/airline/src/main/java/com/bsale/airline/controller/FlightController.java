package com.bsale.airline.controller;

import com.bsale.airline.model.AirlineData;
import com.bsale.airline.model.Flight;
import com.bsale.airline.service.IFlightService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.HttpStatusCodeException;

import java.util.ArrayList;
import java.util.LinkedHashMap;
import java.util.Optional;

@RestController
@RequestMapping //(path="api")
public class FlightController {

    @Autowired
    private IFlightService serviceFlight;

    @GetMapping("/flights/{id}")
    public Optional<Flight> findById(@PathVariable("id") Integer id) {
        return serviceFlight.findById(id);
    }

    @GetMapping("/flights/{id}/passengers")
    public AirlineData data(@PathVariable("id") Integer Id) {
        Optional<Flight> flight =null;
        Integer code=200;
        String message="";

        try {
            flight = serviceFlight.findById(Id);
           if(flight .isEmpty()){
               message="";
               code=404;
           }
        } catch (HttpStatusCodeException e) {
            code=e.getRawStatusCode();
            if(code==400){message="could not connect to db"; }
        }

        return new AirlineData(code, flight,message);
    }


}
