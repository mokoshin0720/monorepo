import type { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from 'App';
import React, { useState } from 'react';
import { View, StyleSheet, TextInput } from 'react-native';
import { KeyboardAvoidingView } from 'react-native';
import { CircleButton } from 'src/components/CircleButton';
import { baseColor } from 'src/utils/color';

type MemoCreateScreenNavigationProp = NativeStackScreenProps<
    RootStackParamList,
    'MemoCreate'
>;

export const MemoCreateScreen: React.FC<MemoCreateScreenNavigationProp> = ({
    navigation,
    route,
}) => {
    const [bodyText, setBodyText] = useState<string>('');

    return (
        <KeyboardAvoidingView style={styles.container} behavior="height">
            <View style={styles.inputContainer}>
                <TextInput
                    value={bodyText}
                    multiline
                    style={styles.input}
                    onChangeText={(text: string) => setBodyText(text)}
                    autoFocus
                />
            </View>
            <CircleButton
                name="check"
                onPress={() => {
                    createMemo(bodyText);
                }}
            />
        </KeyboardAvoidingView>
    );
};

const createMemo = (text: string) => {
    console.log('created!', text);
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: baseColor,
    },
    inputContainer: {
        flex: 1,
        paddingHorizontal: 27,
        paddingVertical: 32,
    },
    input: {
        flex: 1,
        textAlignVertical: 'top',
        fontSize: 16,
        lineHeight: 24,
    },
});
