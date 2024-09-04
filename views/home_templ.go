// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "andrenormanlang/tarot-go-htmx/common"

func Home(cards []common.Card, selectedCards []common.Card, meanings []string, isShuffling bool, revealIndex int) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Free Tarot Reading</title><link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css\" rel=\"stylesheet\"><link href=\"/static/css/styles.css\" rel=\"stylesheet\"><link href=\"https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css\" rel=\"stylesheet\"></head><body class=\"bg-gray-900 text-gray-100\"><div class=\"container mx-auto max-w-4xl px-4 py-8 bg-gray-900\"><h1 class=\"text-2xl sm:text-4xl md:text-5xl lg:text-6xl font-bold mb-6 text-center text-purple-300\">Rider Waite Smith (Mind, Body & Spirit)</h1><!-- Button to trigger the modal --><button type=\"button\" class=\"smythe-regular w-64 text-xl sm:text-2xl bg-purple-600 hover:bg-purple-700 text-white px-6 py-3 rounded-lg mb-6 mx-auto block transition duration-300 ease-in-out\" data-bs-toggle=\"modal\" data-bs-target=\"#infoModal\">Meaning of Card Positions?</button><!-- Do Another Reading Button -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(selectedCards) == 3 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button id=\"doAnotherReading\" hx-get=\"/reset-reading\" hx-target=\"body\" hx-swap=\"innerHTML\" class=\"smythe-regular w-64 mt-12 mb-4 text-xl sm:text-2xl bg-green-600 hover:bg-green-700 text-white px-6 py-3 rounded-lg mx-auto block transition duration-300 ease-in-out\">Do Another Reading</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-full max-w-lg mx-auto mb-8\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(selectedCards) == 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button id=\"shuffleButton\" class=\"smythe-regular w-64 text-xl sm:text-2xl bg-indigo-600 hover:bg-indigo-700 text-white px-6 py-3 rounded-lg mb-4 mx-auto block transition duration-300 ease-in-out\">Shuffle Cards</button> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"text\" placeholder=\"Enter your question or subject here (optional)\" id=\"tarot-question\" class=\"w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-purple-600\"></div><!-- Modal for Card Positions --><div class=\"modal fade\" id=\"infoModal\" tabindex=\"-1\" aria-labelledby=\"infoModalLabel\" aria-hidden=\"true\"><div class=\"modal-dialog modal-lg modal-dialog-centered\"><div class=\"modal-content bg-gray-900 text-purple-400\"><div class=\"modal-header border-b border-gray-700\"><h5 class=\"text-3xl modal-title smythe-regular w-100 text-center\" id=\"infoModalLabel\">Card Positions in the \"Mind, Body & Spirit\" Tarot Spread</h5><button type=\"button\" class=\"btn-close btn-close-white\" data-bs-dismiss=\"modal\" aria-label=\"Close\"></button></div><div class=\"modal-body px-6 py-4\"><table class=\"w-full border-collapse\"><thead><tr class=\"border-b border-gray-700\"><th class=\"text-2xl p-4 text-left smythe-regular\">Position</th><th class=\"text-2xl p-4 text-center smythe-regular\">Meaning</th></tr></thead> <tbody><tr class=\"border-b border-gray-700 hover:bg-gray-800 transition-colors duration-200\"><td class=\"text-xl p-4 font-semibold smythe-regular\">MIND:</td><td class=\"text-xl p-4 text-center old-standard-tt-regular\">The first card represents logic, reasoning, and your rational side. What are you thinking? Or: What is an area of reason you should pay attention to?</td></tr><tr class=\"border-b border-gray-700 hover:bg-gray-800 transition-colors duration-200\"><td class=\"text-xl p-4 font-semibold smythe-regular\">BODY:</td><td class=\"text-xl p-4 text-center old-standard-tt-regular\">The second card represents health, fitness, sexuality, and activity. What are you doing? Or: What is an area of action you should think about?</td></tr><tr class=\"hover:bg-gray-800 transition-colors duration-200\"><td class=\"text-xl p-4 font-semibold smythe-regular\">SPIRIT:</td><td class=\"text-xl p-4 text-center old-standard-tt-regular\">The third and last card represents emotion, creativity, spirituality, and love. What are you feeling? Or: What should you focus on for spiritual growth?</td></tr></tbody></table></div></div></div></div><!-- Shuffled Cards Container --><div id=\"shuffled-cards\" class=\"relative h-36 sm:h-48 mt-2 flex justify-center items-center overflow-x-auto mx-auto\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, card := range cards {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"card-container\" style=\"--card-index: {index};\"><button hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs("/select-card?card=" + card.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home.templ`, Line: 76, Col: 56}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"body\" class=\"bg-transparent\" onclick=\"openModal(&#39;{ card.Name }&#39;, &#39;{ card.MeaningUp }&#39;)\"><img src=\"https://res.cloudinary.com/dytiufsuu/image/upload/v1725446781/rws_tarot/CardBacks_ovn82s.png\" alt=\"Card Back\" class=\"w-24 sm:w-32 h-auto\"></button></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><!-- Selected Cards Container --><div id=\"selected-cards\" class=\"mt-12 flex flex-col md:flex-row justify-center items-center gap-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for i := 0; i < 3; i++ {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col items-center m-4\"><h3 class=\"smythe-regular text-xl sm:text-2xl md:text-3xl font-bold mb-4 text-indigo-300\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if i == 0 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("Past")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if i == 1 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("Present")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("Future")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h3>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if i < len(selectedCards) {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Render the selected card --> <div class=\"flip-card\"><div class=\"flip-card-inner\"><div class=\"flip-card-front\"><img src=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(selectedCards[i].Image)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home.templ`, Line: 100, Col: 44}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" alt=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(selectedCards[i].Name)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home.templ`, Line: 100, Col: 74}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"w-full h-full object-cover rounded-lg\"></div><div class=\"flip-card-back\"><h2 class=\"text-xl sm:text-2xl font-semibold mb-2\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(selectedCards[i].Name)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home.templ`, Line: 103, Col: 85}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><p class=\"text-sm sm:text-base overflow-y-auto flex-grow\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(selectedCards[i].MeaningUp)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home.templ`, Line: 104, Col: 97}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><button type=\"button\" class=\"mt-4 bg-purple-600 hover:bg-purple-700 text-white px-4 py-2 rounded-lg transition duration-300 ease-in-out\" hx-get=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string
				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs("/card-detail?card=" + selectedCards[i].Name)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home.templ`, Line: 108, Col: 65}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#cardModal .modal-content\" data-bs-toggle=\"modal\" data-bs-target=\"#cardModal\">View Full Description</button></div></div></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Render the placeholder --> <div class=\"flip-card\"><div class=\"flip-card-inner\"><div class=\"flip-card-front\"><img src=\"/images/question-mark.svg\" alt=\"Placeholder\" class=\"w-full h-full object-cover rounded-lg\"></div></div></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><!-- Modal for Extended Card Description --><div class=\"modal fade\" id=\"cardModal\" tabindex=\"-1\" role=\"dialog\" aria-labelledby=\"cardModalLabel\" aria-hidden=\"true\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CardDetailModal(common.Card{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><!-- External Scripts --><script src=\"https://cdn.jsdelivr.net/npm/@popperjs/core@2.11.7/dist/umd/popper.min.js\"></script><script src=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.min.js\"></script><script src=\"/static/script/htmx.min.js\"></script><script src=\"/static/script/script.js\"></script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
