Feature: Exclude Files and Folders from AsciiDoc Documentation Generation

  As a user of the source2adoc CLI tool, I want to use a flag to specify files
  and folders that should be ignored so that I can exclude certain files and
  directories from the documentation generation process.

  Initial implementation: https://github.com/sommerfeld-io/source2adoc/issues/102

  Background:
    Given I am using the root command of the source2adoc CLI tool to generate AsciiDoc files

  Scenario: Full Path File Exclusion
    Given I want to exclude a specific file by its full path
    When I specify the filesystem location using the --exclude flag
    Then asciidoc files should be generated for all source code files
      But the tool should not generate asciidoc files for the specified file

  Scenario: Folder Exclusion
    Given I want to exclude an entire folder
    When I specify the filesystem location using the --exclude flag
    Then asciidoc files should be generated for all source code files
    But the tool should not generate asciidoc files for all files and subfolders within that folder

  Scenario Outline: Filename Pattern Exclusion
    Given I want to exclude files that match a specific pattern (e.g., *.sh or prefix-*)
    When I specify the filesystem location using the --exclude flag
    Then asciidoc files should be generated for all source code files
    But the tool should not generate asciidoc files for all files matching the pattern in any directory

    Examples:
      | pattern  |
      | *.sh     |
      | prefix-* |

  Scenario Outline: Pattern with Path Exclusion
    Given I want to exclude files that match a specific pattern within a particular directory
    When I specify the filesystem location using the --exclude flag
    Then asciidoc files should be generated for all source code files
    But the tool should not generate asciidoc files for all files matching the pattern in the specified directory

    Examples:
      | pattern           |
      | src/main/*.sh     |
      | src/main/prefix-* |

  Scenario: Flag Usability with Multiple Values
    Given I want to exclude multiple files and folders in a single command run
    When I specify the --exclude flag multiple times with different values
    Then asciidoc files should be generated for all source code files
    But the tool should not generate asciidoc files for all specified files, folders, and patterns
