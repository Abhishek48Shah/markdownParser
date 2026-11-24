import java.io.ByteArrayOutputStream;
import java.io.FileInputStream;
import java.io.IOException;
import java.nio.charset.StandardCharsets;

public class Main {
	public static byte[] readFile(String filename) throws IOException {
		ByteArrayOutputStream storage = new ByteArrayOutputStream();
		byte[] buff = new byte[1024];
		try (FileInputStream fs = new FileInputStream(filename)) {
			int n;
			while ((n = fs.read(buff)) != -1) {
				storage.write(buff, 0, n);
			}
			return storage.toByteArray();
		} catch (IOException e) {
			throw new IOException("read file: " + e.getMessage());
		}
	}

	public static void main(String[] args) {
		try {
			byte[] data = Main.readFile("content.md");
			String str = new String(data, StandardCharsets.UTF_8);
			System.out.println(str);
		} catch (IOException e) {
			System.out.println(e.getMessage());
		}
	}
}
