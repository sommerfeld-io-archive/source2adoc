## Test the application setup and configuration.

title "Application"

control "app-01" do
    impact 0.7
    title "Check existence of mandatory files"
    desc "Check if mandatory files exist and have the correct permissions."

    FILES = %w(
        /usr/bin/source2adoc
    )
    FILES.each do |folder|
        describe file(folder) do
            it { should exist }
            it { should be_file }
            it { should_not be_directory }
            it { should be_owned_by 'source2adoc' }
            it { should be_grouped_into 'source2adoc' }
            it { should be_readable.by('owner') }
            it { should be_writable.by('owner') }
            it { should be_executable.by('owner') }
            it { should_not be_readable.by('group') }
            it { should_not be_writable.by('group') }
            it { should_not be_executable.by('group') }
            it { should_not be_readable.by('others') }
            it { should_not be_writable.by('others') }
            it { should_not be_executable.by('others') }
        end
    end
end
