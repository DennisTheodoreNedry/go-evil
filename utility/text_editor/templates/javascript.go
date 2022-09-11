package templates

const (
	JS_FUNCS = `
<script>

	//
	//
	// Saves the active written data to a local file
	//
	//
	function save_data() {
		var data = document.getElementById("file_value").value;
		var location = document.getElementById("file_path").value;
		
		if (location === ""){
			location = prompt("Please enter a location to save the file");
		}

		if (location === ""){
			window.alert("The file path can't be empty!");
		}else{
			window.Save(data, location); // Sends the data back to GO
			window.alert("Successfully wrote to "+location);
		}

	}

	//
	//
	// Saves the active written data to a local file
	// and compiles the file
	//
	//
	function compile() {
		var base64 = document.getElementById("base64").innerHTML;
		var name = document.getElementById("output").value;
		var target_os = document.getElementById("os").value;
		var target_arch = document.getElementById("arch").value;
		var obf = document.getElementById("obf").value;
		var ext = document.getElementById("ext").value;

		window.Compile(base64, name, target_os, target_arch, obf, ext);  // Sends the data back to GO
		window.alert("Successfully compiled the malware");
	}

</script>
	`
)
