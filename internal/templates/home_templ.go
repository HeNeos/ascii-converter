// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Home(title string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section><h1>Welcome to my ASCII art generator</h1><p>This is a simple website to convert images to ASCII art. It allows you to upload an image and convert it to an ASCII text format.</p><p>At the moment, it supports common image formats like PNG, JPEG, JPG, and WEBP.</p><p>The output is black and white text at the moment. Now you can download as a PNG image too!!!</p><p style=\"color: #ff6600;\">EXPERIMENTAL FEATURE: You can preserve the colour of the original image but:</p><ul><li>It may stress some computers hardware.</li><li>It will take longer to generate the ASCII art.</li><li>When downloading as image it will still be black and white.</li></ul></section><!-- Form to upload an image and generate ASCII --> <form hx-post=\"/convert-to-ascii\" hx-target=\"#output\" enctype=\"multipart/form-data\"><input type=\"file\" id=\"image\" name=\"image\" accept=\"image/*\" alt=\"Upload image for ASCII conversion\" required> <label><input type=\"checkbox\" name=\"preserve_color\"> Preserve original colors</label> <button id=\"clear-button\">Clear Image</button> <button type=\"submit\">Generate ASCII</button></form><!-- Output container to display ASCII art and download button --> <div id=\"output-container\"><div id=\"output\"></div></div><!-- Script to handle the clear button --> <script>\r\n            const image = document.getElementById('image');\r\n            const output = document.getElementById('output');\r\n            const clearButton = document.getElementById('clear-button');\r\n\r\n            // Add an event listener to the clear button\r\n            clearButton.addEventListener('click', () => {\r\n                image.value = ''; // Clear the image input\r\n                output.innerHTML = ''; // Clear the output\r\n            });\r\n        </script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = Base(title).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
