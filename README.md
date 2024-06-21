<p align="center">
  <img src="https://github.com/murat-cileli/dbee/assets/6532000/8f7a7d54-0904-4296-a6bb-6836ee86a095" />
  <br>
  <strong>DBee</strong>
  <br>
  Fast & Minimalistic Database Browser
</p>

### 🐝 Features
- Simple, lightweight, fast!
- Keyboard-centric workflow
- Single executable with no dependencies
- Supports MySQL, MariaDB, and PostgreSQL
- Save connections (except passwords)
- List database tables and views  
- View table structures and browse data  
- Execute SQL queries  
- SQL query history

### 🐝 Screenshots

| Connection 	| Main 	| Results 	|
|------------	|:----:	|:-------:	|
| <img src="https://github.com/murat-cileli/dbee/assets/6532000/38842b6c-54d2-4be9-9c40-f9c9f679a3d9" style="height:140px;" /> | <img src="https://github.com/murat-cileli/dbee/assets/6532000/8f178b6c-8fa1-4b19-819a-d09e924109b3" style="height:140px;" /> | <img src="https://github.com/murat-cileli/dbee/assets/6532000/c0c91f0a-61fb-4510-8b29-b796ed8ab91f" style="height:140px;" /> |

### 🐝 Installation
Visit [Releases](https://github.com/murat-cileli/dbee/releases) section.

**Building From Source**
```console
git clone https://github.com/murat-cileli/dbee.git
cd src
go build .
./dbee
```

### 🐝 Keyboard Shortcuts

**Global**  
<kbd>ESC</kbd> : Quit application

**Connections Page -> Saved Connections Pane**  
<kbd>Alt</kbd> + <kbd>S</kbd> : Focus saved connections pane  
<kbd>1 .. Z</kbd> : Selects saved connection  
<kbd>Enter</kbd> : Apply saved connection  

**Connections Page -> Connect to Server Pane**  
<kbd>Alt</kbd> + <kbd>D</kbd> : Focus connect to server pane  
<kbd>Tab</kbd> : Focus next input field  

**Main Page -> Database Objects Pane**  
<kbd>Alt</kbd> + <kbd>W</kbd> : Focus database objects pane  
<kbd>1 .. Z</kbd> : Selects a table/view  
<kbd>Enter</kbd> : Browse top 5 table/view rows  
<kbd>Ctrl</kbd> + <kbd>Space</kbd>: View table/view structure

**Main Page -> Query Pane**  
<kbd>Alt</kbd> + <kbd>E</kbd> : Focus query pane  
<kbd>Alt</kbd> + <kbd>Enter</kbd> : Execute SQL query  
<kbd>Alt</kbd> + <kbd>Up</kbd> : Go back in query history  
<kbd>Alt</kbd> + <kbd>Down</kbd> : Go forward in query history  
<kbd>Ctrl</kbd> + <kbd>Shift</kbd> + <kbd>V</kbd> : Paste  
<kbd>Ctrl</kbd> + <kbd>Z</kbd> : Undo  

**Main Page -> Results Table**  
<kbd>Alt</kbd> + <kbd>R</kbd> : Focus results table  
<kbd>Arrow Keys</kbd>, <kbd>Home/End</kbd>, <kbd>Page Up/Page Down</kbd> : Navigate table

### 🐝 Notes
- Contributions are welcome.
- Follow me on [LinkedIn](https://www.linkedin.com/in/murat-cileli/)
