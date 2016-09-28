<?
$filename = $argv[1];
$lines = preg_split(":\n:",trim(file_get_contents($filename)));

function snapshotGenerator($i){
	$i = $i - 1;
	return function() use(&$i){
		$i++;
		return <<<EOD
	set shellCommand to "/usr/sbin/screencapture " & theDesktop & "Screen_Shot_$i.png"
	do shell script shellCommand
EOD;
	};
}

$snapshot = snapshotGenerator(0);

$snapshotBlock = <<<EOD
	set theCurrentDate to (time of(current date))
	set shellCommand to "/usr/sbin/screencapture " & theDesktop & "Screen_Shot_" & theCurrentDate & ".png"
	do shell script shellCommand
EOD;

$header = <<<EOD
tell application "Terminal"
	set theDesktop to POSIX path of (path to desktop as string)
	activate
	set frontWindow to window 1
	set currentTab to do script "echo 'Hello World'"
	tell application "System Events"
		keystroke "+" using {command down}
		keystroke "+" using {command down}
		keystroke "+" using {command down}
		keystroke "+" using {command down}
		keystroke "+" using {command down}
		keystroke "+" using {command down}
		keystroke "+" using {command down}
		keystroke "+" using {command down}
		keystroke "+" using {command down}
		keystroke "f" using {command down, control down}
	end tell
	delay 5
	do script "clear" in currentTab
	delay 10

EOD;

$output = $header;
$output .= $snapshot() . "\n";
foreach ($lines as $line) {
	$output .= "\tset w to do script \"" . $line . "\" in currentTab\n".
		<<<EOD
	repeat 
		delay 1
		if not busy of w then exit repeat
	end repeat

EOD
		. "\tdelay 1\n" 
		. $snapshot() 
		. "\n";
}
$output .= <<<EOD
	delay 5
	tell application "System Events"
		keystroke "w" using {command down}
	end tell
end tell
EOD;
print_r($output);

?>
