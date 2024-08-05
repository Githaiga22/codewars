Complete the solution so that it splits the string into pairs of two characters. If the string contains an odd number of characters then it should replace the missing second character of the final pair with an underscore ('_').
remember to explain in detail please

Examples:
```bash
* 'abc' =>  ['ab', 'c_']
* 'abcdef' => ['ab', 'cd', 'ef']
```
```bash
package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)
var _ = Describe("Split Strings", func() {
    It("Basic tests", func() {
        Expect(Solution("abc")).To(Equal([]string{"ab", "c_"}))
        Expect(Solution("dawsd")).To(Equal([]string{"da", "ws", "d_"}))
        Expect(Solution("awsaws")).To(Equal([]string{"aw", "sa", "ws"}))
    })
})
```