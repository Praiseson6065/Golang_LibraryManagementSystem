#!/bin/bash

psql -h localhost -U lib -p 5432 << EOF

select id,name,email,usertype from users;

EOF
echo "Enter id to change user type to admin:"
unset ids 
read ids
psql -h localhost -U lib -p 5432 << EOF

update users set usertype='admin' where id=$ids ;

EOF


 

