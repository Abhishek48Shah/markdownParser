package parser;

import java.util.ArrayList;

public class Token {
	String type;
	String value;

	public enum Type {
		Heading
	}

	public Token(String type, String value) {
		this.type = type;
		this.value = value;
	}

	private record HeadingInfo(Token tok, int index, boolean flag) {
	}

	private HeadingInfo checkHeading(byte[] markdown, int index, int length) {
		if (index < length && markdown[index] == '#') {
			int count = 0;
			while (index < length && markdown[index] == '#') {
				count++;
				index++;
			}
			if (index < length && markdown[index] == ' ' && count <= 6) {
				String hash = "#";
				var token = new Token("Heading", hash.repeat(count));
				return new HeadingInfo(token, index + 1, true);
			}
		}
		return new HeadingInfo(null, index, false);
	}

	protected ArrayList<Token> createToken(byte[] markdown) {
		int i = 0;
		boolean isAtFirst = true;
		int n = markdown.length;
		var tokens = new ArrayList<Token>();
		while (i < n) {
			if (isAtFirst) {
				HeadingInfo result = checkHeading(markdown, i, n);
				if (result.flag) {
					tokens.add(result.tok);
					i = result.index;
					isAtFirst = false;
					continue;
				}
			}
		}
		return tokens;
	}
}
