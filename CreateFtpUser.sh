useradd -m $1 -g ftp
(echo "$2"; sleep 1; echo "$2";) | passwd $1
