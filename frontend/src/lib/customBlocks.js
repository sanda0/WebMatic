import * as Blockly from "blockly";


//Control blocks
Blockly.Blocks["control_start"] = {
	init: function () {
		this.appendStatementInput("elements")
			.setCheck(null)
			.appendField("Base URL")
			.appendField(new Blockly.FieldTextInput("BaseURL"), "base_url");
		this.setColour(120);
		this.setTooltip("Start");
		this.setHelpUrl("");
	},
};


Blockly.Blocks['control_navigate'] = {
  init: function() {
    this.appendEndRowInput()
        .appendField("Navigate")
        .appendField(new Blockly.FieldTextInput("https://...."), "url");
    this.setPreviousStatement(true, null);
    this.setNextStatement(true, null);
    this.setColour(120);
 this.setTooltip("Navigate");
 this.setHelpUrl("");
  }
};

//element blocks
Blockly.Blocks["element_by_css_selector"] = {
	init: function () {
		this.appendStatementInput("actions")
			.setCheck(null)
			.appendField("CSS Selector")
			.appendField(new Blockly.FieldTextInput("hmtl"), "css_selector");
		this.setPreviousStatement(true, null);
		this.setNextStatement(true, null);
		this.setColour(230);
		this.setTooltip("CSS selector");
		this.setHelpUrl("");
	},
};

Blockly.Blocks["element_by_xpath_selector"] = {
	init: function () {
		this.appendStatementInput("actions")
			.setCheck(null)
			.appendField("XPath Selector")
			.appendField(new Blockly.FieldTextInput("//xpath//"), "xpath_selector");
		this.setPreviousStatement(true, null);
		this.setNextStatement(true, null);
		this.setColour(230);
		this.setTooltip("XPath Selector");
		this.setHelpUrl("");
	},
};

//Action blocks
Blockly.Blocks['action_wait'] = {
  init: function() {
    this.appendEndRowInput()
        .appendField("Wait")
        .appendField(new Blockly.FieldNumber(0, 0, 60), "seconds")
        .appendField("s");
    this.setPreviousStatement(true, null);
    this.setNextStatement(true, null);
    this.setColour(285);
 this.setTooltip("Wait");
 this.setHelpUrl("");
  }
};


Blockly.Blocks['action_click'] = {
  init: function() {
    this.appendEndRowInput()
        .setAlign(Blockly.ALIGN_RIGHT)
        .appendField("Click");
    this.setPreviousStatement(true, null);
    this.setNextStatement(true, null);
    this.setColour(285);
 this.setTooltip("Click");
 this.setHelpUrl("");
  }
};

Blockly.Blocks['action_write'] = {
  init: function() {
    this.appendEndRowInput()
        .appendField("Write")
        .appendField(new Blockly.FieldTextInput("hello..."), "txt");
    this.setInputsInline(true);
    this.setPreviousStatement(true, null);
    this.setNextStatement(true, null);
    this.setColour(285);
 this.setTooltip("Write");
 this.setHelpUrl("");
  }
};

