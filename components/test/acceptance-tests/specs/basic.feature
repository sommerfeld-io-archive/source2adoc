Feature: Generate AsciidDoc Documentation from Source Code files

  As a user of the source2adoc CLI tool, I want to generate AsciiDoc files from source code
  files so that I can create comprehensive and well-structured documentation directly from
  inline comments in the source code.

  Background:
    Given I am using the root command of the source2adoc CLI tool to generate AsciiDoc files

  Scenario Outline: Generate AsciiDoc from Source Code Files
    Given I specify "<path>" using the --source-dir flag
    And I specify "target" using the --output-dir flag
    When I run the app
    Then AsciiDoc files should be generated for all source code files
    # And the path of the source code file should be preserved in the --output-dir directory
    # And the caption of the documentation file should automatically be set from the source code file's name
    # And the path of the source code file should be included in the generated docs file file

    Examples:
      | path          |
      | testdata/good |
      # TODO use a correct path ... remember, bdd tests are not located in components/app

  # TODO: add negative scenario "i expect not to generate docs for ..."
