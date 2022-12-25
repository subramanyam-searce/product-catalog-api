# sudo service postgresql start

# sudo -u postgres psql

# \l
# \c

var1=$(
    echo "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';
    " | sudo -u postgres psql -d banking 
)

echo $var1