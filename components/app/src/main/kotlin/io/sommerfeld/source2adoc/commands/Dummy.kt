package io.sommerfeld.source2adoc.commands

import org.springframework.boot.SpringApplication
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.shell.standard.ShellComponent
import org.springframework.shell.standard.ShellMethod

@ShellComponent
class Dummy {

	@ShellMethod("Add two integers together.")
    fun add(a: Int, b: Int): Int {
        return a + b
    }

    @ShellMethod("Sumstract two integers together.")
    fun sub(a: Int, b: Int): Int {
        return a - b
    }
}
