useradd -m $1
(echo "$2"; sleep 1; echo "$2";) | passwd $1
(echo "$2"; sleep 1; echo "$2" ) | sudo smbpasswd -s -a $1
