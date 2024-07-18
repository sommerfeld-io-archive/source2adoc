package io.sommerfeld.source2adoc.shell

import org.springframework.shell.standard.ShellComponent
import org.springframework.shell.standard.ShellMethod

@ShellComponent
public class Calculator {

    @ShellMethod("Add two integers together")
    fun plus(a: Int, b: Int): Int {
        return a + b
    }

    @ShellMethod("Subtract one integer from another")
    fun minus(a: Int, b: Int): Int {
        return a - b
    }
}
