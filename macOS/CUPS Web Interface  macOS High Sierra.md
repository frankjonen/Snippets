# CUPS Web Interface: macOS High Sierra

To get an overview of your print spooler you can enable the web interface for CUPS. 

> (this needs admin privileges, so login as an admin user first)
> `login admin` for example.

```bash
sudo cupsctl WebInterface=yes
```

You can reach it via `http://localhost:631`

To view your printers: `http://localhost:631/printers/`
