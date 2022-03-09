<!DOCTYPE html>
<html>
<head>
    <meta charset='utf-8'>
    <title>Aufgabe hinzufügen</title>
    <meta name='viewport' content='width=device-width, initial-scale=1'>
    <link rel='stylesheet' type='text/css' media='screen' href='main.css'>
    <script src='main.js'></script>
</head>
<body>
    <ul>
        <li><a href="index.html">Zurück</a></li>
    </ul>
    <h1 class="text-glow">Füge eine Aufgabe hinzu</h1><br>
    <form style="text-align: center;" action="submit.php" method="post" id="task-input">
        <input type="text" id="task-input-field" name="input-field" class="task-input" value="LOL"><br><br>
        <button type="submit" id="submit-button" class="button-form">Hinzufügen</button>
    </form>
</body>
</html>