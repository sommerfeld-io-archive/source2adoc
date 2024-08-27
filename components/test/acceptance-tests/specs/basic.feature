Feature: Generate AsciidDoc Documentation from Source Code files

  As a user of the source2adoc CLI tool, I want to generate AsciiDoc files from source code
  files so that I can create comprehensive and well-structured documentation directly from
  inline comments in the source code.

  Background:
    Given I am using the root command of the source2adoc CLI tool to generate AsciiDoc files

  Scenario: Display help message
    Given I specify the "--help" flag
    When I run the app
    Then exit code should be 0

  Scenario Outline: Generate AsciiDoc for supported Source Code Files
    Given I specify the "--source-dir" flag with value "<path>"
    And I specify the "--output-dir" flag with value "/workspaces/source2adoc/target"
    When I run the app with volume mount "/workspaces/source2adoc"
    Then exit code should be 0
    # And AsciiDoc files should be generated for all source code files
    # And the path of the source code file should be preserved in the --output-dir directory
    # And the caption of the documentation file should automatically be set from the source code file's name
    # And the path of the source code file should be included in the generated docs file file

    Examples:
      | path |
      | /workspaces/source2adoc/testdata/common/good |

  Scenario Outline: No AsciiDoc for unsupported Source Code Files
    Given I specify the "--source-dir" flag with value "<path>"
    And I specify the "--output-dir" flag with value "/workspaces/source2adoc/target"
    When I run the app with volume mount "/workspaces/source2adoc"
    Then exit code should be 0
    # And no AsciiDoc files should be generated

    Examples:
      | path |
      | /workspaces/source2adoc/testdata/common/bad |

  Scenario Outline: Error message for missing source dir
    Given I specify the "--source-dir" flag with value "<path>"
    And I specify the "--output-dir" flag with value "/workspaces/source2adoc/target"
    When I run the app with volume mount "/workspaces/source2adoc"
    Then exit code should be 1
    # And no AsciiDoc files should be generated

    Examples:
      | path |
      | /workspaces/source2adoc/testdata/missing |

  # Scenario: Missing mandatory flags
  #  Given ...
  #  ...
  #  Then exit code should be 1
