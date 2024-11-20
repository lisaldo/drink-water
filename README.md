# Drink-water
A simple reminder to drink water

## Creating daemon
[Doc](https://wiki.archlinux.org/title/systemd/User#Automatic_start-up_of_systemd_user_instances)

```
go build cmd/main.go;
mkdir -p $HOME/.local/bin;
mv main $HOME/.local/bin/drink-water;
mkdir -p $HOME/.config/systemd/user;
cat << EOF > $HOME/.config/systemd/user/drink-water.service
[Unit]
Description=Reminder para beber agua
DefaultDependencies=no

[Service]
ExecStart=$HOME/.local/bin/drink-water
Type=simple

[Install]
WantedBy=default.target

EOF
loginctl enable-linger;
systemctl --user daemon-reload;
systemctl --user start drink-water.service;
```

