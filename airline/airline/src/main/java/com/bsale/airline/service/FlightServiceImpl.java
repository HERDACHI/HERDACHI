package com.bsale.airline.service;

import com.bsale.airline.model.Flight;
import com.bsale.airline.repository.IFlightRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Service;

import java.util.Optional;

@Service("FlightService")
public class FlightServiceImpl implements IFlightService {

    @Autowired
    @Qualifier("FlightRepository")
    private IFlightRepository crud;

    @Override
    public Optional<Flight> findById(Integer Id) {
        return crud.findById(Id);
    }

    @Override
    public Flight buscarVuelo(Integer Id) {
        return (Flight) crud.buscarVuelo(Id);
    }

}