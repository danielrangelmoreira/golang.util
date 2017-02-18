package main

import (
	//"bufio"
	//"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

//Foram gerados 5 parágrafos, 471 palavras e 3177 bytes de Lorem Ipsum

const lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras suscipit convallis ipsum eget sodales. Praesent a volutpat felis. Fusce tristique elit vitae erat rutrum elementum. Duis consectetur auctor diam vel commodo. Aenean tempus velit in pellentesque auctor. Sed at nisl blandit, pretium magna suscipit, efficitur arcu. Nulla mattis interdum ipsum, at rutrum elit ultricies vitae. Sed sollicitudin libero non velit vestibulum semper. Maecenas imperdiet ipsum eros, eu dapibus sem ultrices ut.
Vivamus in accumsan ex. Curabitur ex risus, ultricies quis rhoncus ornare, sollicitudin ac nunc. Sed consectetur, quam in aliquam lacinia, metus quam luctus ligula, ut dictum lectus dui eu odio. In sodales tempor lorem, vitae ullamcorper purus imperdiet vel. Nulla vestibulum molestie velit elementum tristique. Etiam vel tempus arcu, et elementum neque. Mauris vel ullamcorper dolor. Nullam interdum turpis ut libero condimentum, nec commodo enim pharetra. Proin efficitur mi leo, ac pretium sem volutpat sit amet. Nunc eu eros suscipit, auctor nisl et, finibus risus. Sed auctor imperdiet aliquet. Etiam magna mi, viverra non laoreet nec, volutpat vel orci. Quisque nec diam et arcu accumsan tempus.
In sed est nec ex pharetra vulputate bibendum ut erat. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis diam nulla, cursus id convallis a, maximus id elit. Nunc lacinia lorem vel augue mattis, eu placerat est dignissim. Vestibulum placerat sapien vitae leo tincidunt semper. Praesent viverra convallis blandit. Nunc aliquet enim enim, a eleifend eros tristique sit amet. Nunc sit amet commodo metus. Nullam vitae dolor arcu. Mauris feugiat nunc sit amet rutrum sollicitudin. In a arcu massa. Nunc non justo nec massa cursus faucibus sed nec lectus. Sed erat felis, feugiat non mauris vitae, tristique euismod neque. Sed egestas consectetur sodales.
Donec facilisis neque feugiat libero gravida suscipit. Fusce rutrum est gravida iaculis finibus. Praesent commodo, risus vel auctor elementum, nibh nulla fringilla eros, ut pretium nulla mauris et lectus. Mauris vel egestas elit. Mauris pellentesque placerat suscipit. Aenean finibus cursus dapibus. Sed feugiat convallis orci sed rhoncus. Aenean dignissim orci risus, non tincidunt neque cursus vitae. Morbi aliquam erat sit amet leo dignissim, a dignissim lectus pulvinar. Nam pulvinar dui ipsum, at lacinia elit efficitur ac. Nam convallis aliquam sem, eu porta leo. Suspendisse augue elit, euismod non pretium ac, aliquet vel tortor. Aenean eu odio et mauris tincidunt sodales. Aliquam tempor urna tortor, vel hendrerit ipsum egestas quis. In faucibus scelerisque justo consequat lobortis.
Quisque pretium lectus quis erat congue dictum. Fusce mauris dolor, fermentum ut mi in, convallis volutpat velit. Aenean scelerisque nisi eu lectus finibus egestas. Duis placerat eleifend ante. Etiam rhoncus congue nisi, nec tristique dolor gravida ut. Donec condimentum metus lorem, non efficitur felis efficitur in. Vivamus consectetur lectus purus, ut porta quam egestas quis. Vestibulum ultricies magna pulvinar dui efficitur convallis. Proin vel suscipit arcu. Donec tempor est ac varius placerat. Fusce ac pretium arcu.`
const lorem2 = `Lorem ipsum dolor sit amet, consectetur`

type CountWrite struct {
	io.Writer
	count int64
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var z = &CountWrite{Writer: w}

	return z.Writer, &z.count
}

func (c *CountWrite) Write(p []byte) (int, error) {
	a, err := c.Writer.Write(p)
	if err != nil {
		return a, fmt.Errorf("CountWriting: %s", err)
	}
	c.count += int64(a)

	return len(p), nil
}

func main() {

	dst := &CountWrite{Writer: os.Stdout}

	fmt.Fprintf(dst, "%s", strings.Repeat(" ", 10000))
	fmt.Println()
	fmt.Println(dst.count)

}
