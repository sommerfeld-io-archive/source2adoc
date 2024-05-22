##
#
##

title 'devcontainer checks'

control 'devcontainer-1.0' do
    impact 0.5
    title "Verify devcontainer operating system"
    desc "Verify the devcontainer operating system type and configuration"

    describe os.family do
        it { should eq 'debian' }
    end
end
