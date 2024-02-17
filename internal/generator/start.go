/*
Copyright 2024 lucasloureiror

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package generator

import (
	"github.com/lucasloureiror/AegisPass/internal/charsets"
	"github.com/lucasloureiror/AegisPass/internal/cli"
)

func Start(input cli.Input, mode PasswordGeneratorStrategy) {

	generator := passwordGenerator{
		mode: mode,
		data: input,
	}

	generator.generated = make([]string, input.NumberOfPasswords)

	charsets.Create(&generator.data)

	for i := 0; i < input.NumberOfPasswords; i++ {
		generator.generated[i], generator.credits, _ = generator.mode.generate(&generator.data)
	}

	generator.print()

}
