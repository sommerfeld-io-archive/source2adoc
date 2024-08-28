## Test the website content and filesytem permissions.

title "Website content"

control "website-01" do
    impact 0.7
    title "Check existence of mandatory folders"
    desc "Check folders containing the webserver configuration and the website content"

    FOLDERS = %w(
        /usr/local/apache2/htdocs
    )
    FOLDERS.each do |folder|
        describe file(folder) do
            it { should exist }
            it { should_not be_file }
            it { should be_directory }
            it { should be_owned_by apache.user }
            it { should be_grouped_into apache.user }
            it { should be_readable.by('owner') }
            it { should be_writable.by('owner') }
            it { should be_executable.by('owner') }
            it { should be_readable.by('group') }
            it { should_not be_writable.by('group') }
            it { should be_executable.by('group') }
            it { should be_readable.by('others') }
            it { should_not be_writable.by('others') }
            it { should be_executable.by('others') }
        end
    end
end

control "website-02" do
    impact 0.7
    title "Check existence of mandatory files"
    desc "Check files containing the webserver configuration and the website content"

    FILES = %w(
        /usr/local/apache2/conf/httpd.conf
        /usr/local/apache2/htdocs/index.html
        /usr/local/apache2/htdocs/robots.txt
    )
    FILES.each do |file|
        describe file(file) do
            it { should exist }
            it { should be_file }
            it { should_not be_directory }
            it { should be_owned_by apache.user }
            it { should be_grouped_into apache.user }
            it { should be_readable.by('owner') }
            it { should be_writable.by('owner') }
            it { should be_readable.by('group') }
            it { should_not be_writable.by('group') }
            it { should be_readable.by('others') }
            it { should_not be_writable.by('others') }
            it { should_not be_executable }
        end
    end
end
