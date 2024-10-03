package com.bsale.airline.model;

import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.Optional;


@JsonInclude(JsonInclude.Include.NON_EMPTY)
public class AirlineData {
    private Integer code;
    private Optional<Flight> flight;

    private String  errors;

    public AirlineData(Integer code, Optional<Flight> flight, String errors) {
        this.code = code;
        this.flight = flight;
        this.errors =errors;
    }

    public String getErrors() {
        return  errors;
    }

    public Integer getCode() {
        return code;
    }

    public Object getData() {
        return flight;
    }

}
