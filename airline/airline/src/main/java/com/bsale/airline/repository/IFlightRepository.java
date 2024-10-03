package com.bsale.airline.repository;

import com.bsale.airline.model.Flight;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

@Repository("FlightRepository")
public interface IFlightRepository extends JpaRepository<Flight, Integer> {
    @Query(value = "SELECT * FROM flight f JOIN boarding_pass bp ON f.flight_id = bp.flight_id JOIN passenger p ON bp.passenger_id= p.passenger_id  WHERE bp.flight_id=?1 ORDER BY bp.passenger_id", nativeQuery = true)
    Flight buscarVuelo(int valor);
}
