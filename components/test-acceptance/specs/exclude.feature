Feature: Exclude Files and Folders from AsciiDoc Documentation Generation

  As a user of the source2adoc CLI tool, I want to use a flag to specify files
  and folders that should be ignored so that I can exclude certain files and
  directories from the documentation generation process.

  Background:
    Given I use the "root" command of the source2adoc CLI tool

  Scenario Outline: Exclue specific file by its full path
    Given I specify the "--source-dir" flag with value "/workspaces/source2adoc/testdata/common"
    And I specify the "--output-dir" flag with value "/workspaces/source2adoc/target/acceptance-test"
    And I specify the "--exclude" flag with value "<exclude>"
    When I run the app
    Then exit code should be 0
    And AsciiDoc files should be generated for all source code files
    # But no AsciiDoc file should be generated for the excluded

    Examples:
      | exclude                                                |
      | /workspaces/source2adoc/testdata/common/good/Makefile  |
      | /workspaces/source2adoc/testdata/common/good/script.sh |

  Scenario Outline: Exclude Entire Folder
    Given I specify the "--source-dir" flag with value "/workspaces/source2adoc/testdata/common"
    And I specify the "--output-dir" flag with value "/workspaces/source2adoc/target/acceptance-test"
    And I specify the "--exclude" flag with value "<exclude>"
    When I run the app
    Then exit code should be 0
    And AsciiDoc files should be generated for all source code files
    # But no AsciiDoc file should be generated for the excluded

    Examples:
      | exclude                                                      |
      | /workspaces/source2adoc/testdata/testdata/common/good/docker |
      | /workspaces/source2adoc/testdata/testdata/common/good/yaml |

  Scenario Outline: Exclude multiple files and folders in a single command run
    Given I specify the "--source-dir" flag with value "/workspaces/source2adoc/testdata/common"
    And I specify the "--output-dir" flag with value "/workspaces/source2adoc/target/acceptance-test"
    And I specify the "--exclude" flag with value "<exclude-1>"
    And I specify the "--exclude" flag with value "<exclude-2>"
    When I run the app
    Then exit code should be 0
    And AsciiDoc files should be generated for all source code files
    # But no AsciiDoc file should be generated for the excluded

    Examples:
      | exclude-1 | exclude-2 |
      | Makefile  | script.sh |
      | docker    | yaml      |
      | docker    | script.sh |

  Scenario: Exclude flag without value
    Given I specify the "--source-dir" flag with value "/workspaces/source2adoc/testdata/common"
    And I specify the "--output-dir" flag with value "/workspaces/source2adoc/target/acceptance-test"
    And I specify the "--exclude" flag
    When I run the app
    Then exit code should be 1
