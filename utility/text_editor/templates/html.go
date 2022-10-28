package templates

const (
	HTML = `
	%s
	%s
	<body>
		<h1>%s</h1>
		<h5>Version: %s</h5>
		<div class="parent">
		
			<div class="text_box">
				<textarea type="text" id="file_value">%s</textarea>
				<br/>
				<input type="submit" value="Save" onclick="save_data()">
				<input type="submit" value="Compile" onclick="compile()">
			</div>

			<div class="options">
				<h2>Options</h2>

				<label for="file_path">File location *</label>
				<input type="text" id="file_path" value="%s">
				
				<br/><br/>
				<label for="output">Output</label>
				<input type="text" id="output" placeholder="me_no_virus">
				
				<br/><br/>
				<label for="os">Target OS</label>
				<select id="os">
					<option value="default">Default</option>
					<option value="windows">Windows</option>
					<option value="linux">Linux</option>
					<option value="darwin">Darwin</option>
			  	</select>

				<br/><br/>
				<label for="arch">Target architecture</label>
				<select id="arch">
					<option value="default">Default</option>
					<option value="amd64">amd64</option>
					<option value="i386">i386</option>
				</select>
			
				<br/><br/>
				<label for="obf">Obfusicate code</label>
				<select id="obf">
					<option value="yes">Yes</option>
					<option value="no">No</option>
			 	 </select>

			  	<br/><br/>
				<label for="ext">Extension</label>
				<input type="text" id="ext" placeholder="">
			</div>

			<br/><br/>
		</div>

		<p class="base64" id="base64">%s</p>
	</body>`
)
