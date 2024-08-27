Feature: Exclude Files and Folders from AsciiDoc Documentation Generation

  As a user of the source2adoc CLI tool, I want to use a flag to specify files
  and folders that should be ignored so that I can exclude certain files and
  directories from the documentation generation process.

  Background:
    Given I am using the root command of the source2adoc CLI tool to generate AsciiDoc files

  Scenario Outline: Exclue specific file by its full path
    Given I specify "<path>" using the --exclude flag
    When I run the app
    Then exit code should be 0
    And AsciiDoc files should be generated for all source code files
    But the tool should not generate AsciiDoc files for the specified file

    Examples:
      | path                           |
      | testdata/common/good/Makefile  |
      | testdata/common/good/script.sh |
      # TODO use a correct path ... remember, bdd tests are not located in components/app

  Scenario Outline: Exclude Entire Folder
    Given I specify "<path>" using the --exclude flag
    When I run the app
    Then exit code should be 0
    And AsciiDoc files should be generated for all source code files
    But the tool should not generate AsciiDoc files for all files and subfolders within that folder

    Examples:
      | path                        |
      | testdata/common/good/docker |
      # TODO use a correct path ... remember, bdd tests are not located in components/app

  # TODO https://github.com/sommerfeld-io/source2adoc/issues/109
  # Scenario Outline: Exclude files or folders that match a specific pattern
  #   Given I specify "<pattern>" using the --exclude flag
  #   When I run the app
  #   Then exit code should be 0
  #   And AsciiDoc files should be generated for all source code files
  #   But the tool should not generate AsciiDoc files for all files matching the pattern in any directory
  #
  #   Examples:
  #     | pattern  |
  #     | *.sh     |
  #     | prefix-* |

  # TODO https://github.com/sommerfeld-io/source2adoc/issues/109
  # Scenario Outline: Exclude files or folders that match a specific pattern within a particular directory
  #   Given I specify "<pattern>" using the --exclude flag
  #   When I run the app
  #   Then exit code should be 0
  #   And AsciiDoc files should be generated for all source code files
  #   But the tool should not generate AsciiDoc files for all files matching the pattern in the specified directory
  #
  #   Examples:
  #     | pattern                  |
  #     | src/main/*.sh            |
  #     | src/main/prefix-*        |
  #     | src/main/**/sources      |
  #     | src/main/**/sources/*.sh |

  Scenario: Exclude multiple files and folders in a single command run
    Given I specify the --exclude flag multiple times with different values
    When I run the app
    Then exit code should be 0
    And AsciiDoc files should be generated for all source code files
    But the tool should not generate AsciiDoc files for all specified files, folders, and patterns
    # TODO concretize ... test is not deterministic enough ... to vague
