package io.sommerfeld.source2adoc.shell

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

import io.sommerfeld.source2adoc.shell.Calculator

class CalculatorTest {

    private val calculator = Calculator()

    @Test
    fun testPlus() {
        val result = calculator.plus(2, 3)
        assertEquals(5, result)
    }

    @Test
    fun testMinus() {
        val result = calculator.minus(5, 2)
        assertEquals(3, result)
    }
}
