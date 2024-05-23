## This InSpec test aims to validate basic operating system settings on the target system. It covers
## a range of basic configurations to ensure compliance, security, and expected behavior.
##
## Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt
## ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo
## dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit
## amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor
## invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et
## justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum
## dolor sit amet.
##
## CAUTION: This is a file from the test data set

title 'check basic system configuration'

## Do I have to change the language from ruby to chef inspec or something?
control 'basic-01' do
    impact 1.0
    title 'Check basic system configuration'
    desc 'Ensure tests are run against the correct machine and ensure basic configuration is correct'

    describe os.family do
        it { should cmp 'debian' }
    end

    describe timezone do
        its('identifier') { should cmp 'Europe/Berlin' }
    end
end
