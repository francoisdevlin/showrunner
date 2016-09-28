<?
$filename = $argv[1];
$lines = preg_split(":\n:",trim(file_get_contents($filename)));
//$lines = [
	//"echo 'Bacon'",
	//"pwd",
	//"ls"
	//];
$snapshotBlock = <<<EOD
	set theCurrentDate to (time of(current date))
	set shellCommand to "/usr/sbin/screencapture " & theDesktop & "Screen_Shot_" & theCurrentDate & ".png"
	do shell script shellCommand
EOD;

$header = <<<EOD
tell application "Terminal"
	set theDesktop to POSIX path of (path to desktop as string)
	activate
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
foreach ($lines as $line) {
	$output .= "\tdo script \"" . $line . "\" in currentTab\n\tdelay 5\n" . $snapshotBlock . "\n";
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
