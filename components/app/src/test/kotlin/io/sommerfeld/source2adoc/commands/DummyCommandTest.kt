package io.sommerfeld.source2adoc.commands

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import DummyCommand

class DummyCommandTest {
    private val dummy = DummyCommand()

    @Test
    fun testAdd() {
        val result = dummy.add(2, 3)
        assertEquals(5, result)
    }

    @Test
    fun testSub() {
        val result = dummy.sub(5, 2)
        assertEquals(3, result)
    }
}
