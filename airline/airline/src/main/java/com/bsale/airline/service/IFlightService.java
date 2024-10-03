package com.bsale.airline.service;

import com.bsale.airline.model.Flight;

import java.util.Optional;

public interface IFlightService {

    Optional<Flight> findById(Integer Id);

    Flight buscarVuelo(Integer valor);

}
