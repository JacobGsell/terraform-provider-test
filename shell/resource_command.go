package shell

import (
	"os/exec"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCommand() *schema.Resource {
	return &schema.Resource{
		Create: resourceCommandCreate,
		Read:   resourceCommandRead,
		Update: resourceCommandUpdate,
		Delete: resourceCommandDelete,

		Schema: map[string]*schema.Schema{
			"command": {
				Type:     schema.TypeString,
				Required: true,
			},
			"output": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceCommandCreate(d *schema.ResourceData, m interface{}) error {
	return resourceCommandRun(d)
}

func resourceCommandRead(d *schema.ResourceData, m interface{}) error {
	// This resource does not have a state to read, so we just return nil
	return nil
}

func resourceCommandUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceCommandRun(d)
}

func resourceCommandDelete(d *schema.ResourceData, m interface{}) error {
	// Nothing to delete, so just remove from state
	d.SetId("")
	return nil
}
func resourceCommandRun(d *schema.ResourceData) error {
	cmd := d.Get("command").(string)

	// Windows-spezifische Anpassung für den cmd-Interpreter
	// Hier wird "cmd.exe" direkt verwendet, und "/C" wird verwendet, um den Befehl auszuführen und die Eingabeaufforderung danach zu schließen.
	out, err := exec.Command("powershell.exe", "/C", cmd).CombinedOutput()
	if err != nil {
		return err
	}

	// Setze die Ausgabe und die ID für den ResourceData
	d.Set("output", "the output"+string(out))
	d.SetId(cmd)

	return nil
}
