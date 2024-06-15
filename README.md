<p align="center">
  <img src="https://github.com/murat-cileli/dbee/blob/main/dbee-logo.png" />
  <br>
  <strong>DBee - TUI Database Browser</strong>
</p>

DBee is a terminal-based TUI application written in Go, designed for connecting to MySQL, MariaDB, and PostgreSQL databases. It offers a terminal-based interface for browsing database contents and executing SQL queries.

### Features
- Simple, lightweight, fast!
- Single executable with no dependencies
- Supports MySQL, MariaDB, and PostgreSQL
- Keyboard-centric workflow
- Save connections (except passwords)
- List database tables and views
- View table structures and browse data
- Execute SQL queries
- Fixed columns in results table  
  
[](https://github.com/murat-cileli/dbee/assets/6532000/d9d2cd86-e505-471d-91e4-d56cf8d34725)  

## Keyboard Shortcuts

**Connections Page - Saved Connections Pane**  
<kbd>Alt</kbd> + <kbd>S</kbd> : Focus saved connections panel  
<kbd>1 .. Z</kbd> : Selects saved connection  
<kbd>Enter</kbd> : Apply saved connection  

**Connections Page - Connect to Server Pane**  
<kbd>Alt</kbd> + <kbd>D</kbd> : Focus connection details panel  
<kbd>Tab</kbd> : Focus next input field (on connection details panel)  

**Main Page - Database Objects Pane**  
<kbd>Alt</kbd> + <kbd>W</kbd> : Focus database objects panel  
<kbd>1 .. Z</kbd> : Selects an object  
<kbd>Enter</kbd> : Browse top 10 database object rows  
<kbd>Ctrl</kbd> + <kbd>Space</kbd>: Browse top 10 database object rows  

**Main Page - Query Pane**  
<kbd>Alt</kbd> + <kbd>Q</kbd> : Focus query pane  
<kbd>Alt</kbd> + <kbd>Enter</kbd> : Execute SQL query  
<kbd>Ctrl</kbd> + <kbd>Shift</kbd> + <kbd>V</kbd> : Paste  
<kbd>Ctrl</kbd> + <kbd>Z</kbd> : Undo  

**Main Page - Results Table**  
<kbd>Alt</kbd> + <kbd>R</kbd> : Focus results table  
<kbd>Arrow Keys</kbd>, <kbd>Home/End</kbd>, <kbd>Page Up/Page Down</kbd> : Navigate table

### Download
- [Windows (amd64)](https://github.com/murat-cileli/dbee/releases/download/0.1/dbee_windows_amd64.exe)
- [GNU/Linux (amd64)](https://github.com/murat-cileli/dbee/releases/download/0.1/dbee_linux_amd64)
