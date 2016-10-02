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
	delay 5

EOD;

$output = $header;
$output .= $snapshot() . "\n";
foreach ($lines as $line) {
	if(substr($line,0,1) == "!"){
		$output .= enterKeyStrokes(substr($line,1));
	}else{
		$output .= waitForScript($line);
	}
	$output.= "\tdelay 1\n" 
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
//print_r($output);

$tempFile = "/tmp/temp_command.applescript";
file_put_contents($tempFile,$output);
exec("/usr/bin/osascript $tempFile");
//exec("rm -rf $tempFile");


function waitForScript($line){
	$output = "\tset w to do script \"" . $line . "\" in currentTab\n".
		<<<EOD
	repeat 
		delay 1
		if not busy of w then exit repeat
	end repeat

EOD;
	return $output;
}


function printWord($word){
	$chars = [
		"1"=>"18",
		"2"=>"19",
		"3"=>"20",
		"4"=>"21",
		"5"=>"23",
		"6"=>"22",
		"7"=>"26",
		"8"=>"28",
		"9"=>"25",
		"0"=>"29",
		"+"=>"24 using shift down",
		"-"=>"27",
		];
	$buffer = "";
	$output = "\ttell application \"System Events\"\n";
	$letters = [];
	for($i=0;$i<strlen($word);$i++){
		$letters[]=substr($word,$i,1);
	}
	foreach ($letters as $letter){
		if(array_key_exists($letter,$chars)){
			if($buffer != ""){
				$output .= "\t\tkeystroke \"$buffer\"\n";
				$buffer = "";
			}
			$output .= "\t\tkey code " . $chars[$letter] . "\n";
		}else{
			$buffer .= $letter;
		}
	}
	if($buffer != ""){
		$output .= "\t\tkeystroke \"$buffer\"\n";
		$buffer = "";
	}
	$output .= "\tend tell\n";
	return $output;
}


function enterKeystrokes($line){
	$words = preg_split(":\s+:",$line);
	$output.="";
	foreach ($words as $word){
		//$output.="\ttell application \"System Events\" to keystroke \"$word\"\n";
		$output.=printWord($word);
		$output.="\tdelay .1\n";
		$output.="\ttell application \"System Events\" to keystroke space\n";
		$output.="\tdelay .1\n";
	}
	$output.="\ttell application \"System Events\" to keystroke return\n";
	return $output;
}
?>