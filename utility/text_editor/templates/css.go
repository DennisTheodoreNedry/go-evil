package templates

const (
	CSS = `
<style>
	body{
		background-color: #E6E6E3;
	}

	.parent{
		display: grid;
		grid-template-columns: 2fr 1fr;
		grid-gap: 20px;
	}


	div.text_box textarea {
		width: 700px;
		height: 550px;
		resize: none;
	}

	h1, h2, h5 {
		text-align: center;
	}

	.options {
		width: 80%;
		float: left;
		padding-left: 2%;
		padding-right: 2%;
		border: 2px solid black;	
	}


	div.options input {
		width:100%;
	}

	.base64{
		color: transparent;
	}

</style>
	`
)
